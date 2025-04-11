# Zachary Perry
# COSC 581 Assignment 7: Part 2b
# 4/10/25
import gurobipy as gp
from gurobipy import GRB
import sys


def read_file_build_graph(file_name: str):
    '''
        The read file and build graph function does exactly that. It reads in
        the file containing the graph input and creates the graph.

        The input file cointains a graph in modified DIMACS format, where the
        first line has the total number of vertices and edges and the rest of
        the lines have two values, both vertices representing an edge between
        them.

        Input:
            file_name (str): name of the input file
        Returns:
            graph (2d arr): 2d array representing the graph as an adjacency lst
            vertices (int): total number of vertices read
    '''
    graph = []

    with open(file_name, "r") as file:
        contents = file.readlines()

    vert_edge_line = contents[0].strip().split()
    if len(vert_edge_line) < 2 or len(vert_edge_line) > 2:
        print("error (read_file_build_graph): error parsing first line in the file. Should only include the vertex count and edge count")
        sys.exit(1)
    vertices = int(vert_edge_line[0])
    edges = int(vert_edge_line[1])

    # Initialize graph (adjacency list).
    for i in range(vertices):
        graph.append([])

    edges_read = 0

    # Build adjacency list.
    for i in range(1, len(contents)):
        # NOTE: Ensures we don't read extra edges
        if edges_read > edges:
            break
        curr_edge = contents[i].strip().split()
        if len(curr_edge) > 2 or len(curr_edge) < 2:
            print(
                "error (read_file_build_graph): error parsing edge. Either too little or too few arguments.")
            sys.exit(1)

        vertex_1 = int(curr_edge[0])
        vertex_2 = int(curr_edge[1])

        graph[vertex_1].append(vertex_2)
        graph[vertex_2].append(vertex_1)

        edges_read += 1

    return graph, vertices


def solve_max_clique_ilp(graph, num_vertices):
    '''
        This function will cast the maximum clique problem as an ILP problem.
        It will then solve it using Gurobi.

        Inputs:
            graph (2d arr): 2d array representing the graph as an adjacency lst
            vertices (int): total number of vertices read

        Returns:
            clique: maximum clique found

    '''

    # Create a new model gurobi model to use for the max clique problem
    gurobi_model = gp.Model("Max_Clique")

    # Track the decisions for each vertex
    # (i.e. whether it is in the max clique or not)
    decisions = {}
    for i in range(num_vertices):
        decisions[i] = gurobi_model.addVar(vtype=GRB.BINARY, name=f"x_{i}")

    # Set the object to the maximum clique size
    gurobi_model.setObjective(gp.quicksum(
        decisions[i] for i in range(num_vertices)), GRB.MAXIMIZE)

    # Add a constraint here. Basically, if any pair of vertices are not connected to one another,
    # then only one of those vertices can technically be in the clique.
    for i in range(num_vertices):
        for j in range(i+1, num_vertices):
            if j not in graph[i]:
                gurobi_model.addConstr(
                    decisions[i] + decisions[j] <= 1, f"no_edge_{i}_{j}")

    gurobi_model.optimize()

    # Now, we can get the solution max clique found
    if gurobi_model.status == GRB.OPTIMAL:
        clique = set()
        for i in range(num_vertices):
            # Determining if the decision for the vertex is either in or not in the clique
            if decisions[i].X > 0.5:
                clique.add(i)
        return clique

    return set()


def main():
    if len(sys.argv) < 2:
        print("usage: python3 problem_2a.py filename")
        sys.exit(1)

    file_name = sys.argv[1]

    # Read in the file to build the graph and then find the max clique
    graph, num_vertices = read_file_build_graph(file_name)
    max_clique = solve_max_clique_ilp(graph, num_vertices)

    print("Max clique: ", max_clique)


if __name__ == "__main__":
    main()
