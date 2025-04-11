# Zachary Perry
# COSC 581 Assignment 7: Part 2a
# 4/10/25

import sys
import time


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

    # Get the first line containing the number of vertices and edges
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
            print("error (read_file_build_graph): error parsing edge. Either too little or too few arguments.")
            sys.exit(1)

        vertex_1 = int(curr_edge[0])
        vertex_2 = int(curr_edge[1])

        graph[vertex_1].append(vertex_2)
        graph[vertex_2].append(vertex_1)

        edges_read += 1

    return graph, vertices


def bron_kerbosch(R, P, X, graph, max_clique):
    '''
         Implementation of the standard Bron Kerbosch Algorithm for
         finding the max clique in a graph.

         Source: https://en.wikipedia.org/wiki/Bron–Kerbosch_algorithm

         Input:
            R: The set of vertices currently a part of the potential clique being built
            P: The set of vertices that are all are connected to every vertex in R and could potentially be fully added
            X: The set of vertices that have already been processed
            graph (2d arr): 2d array representing the graph as an adjacency lst
            max_clique: Set containing the current vertices within the max clique found
         Returns:

    '''

    # If both P and X are empty: then report R as the current max clique.
    if not P and not X:
        if len(R) > len(max_clique[0]):
            max_clique[0] = R.copy()
        return

    # Otherwise, recursively explore whether each vertex in P can be added to the current clique.
    # 1. grab all neighbors of the current vertex
    # 2. Recursively call BK with:
        # 2a: R after adding the current vertex to the clique
        # 2b: P now equals the intersection of vertices with the current neighbors of the current vertex
        # 2c: X now equals the intersection of vertices with the current neighbors of the current vertex
    for vertex in list(P):
        curr_neighbors = set(graph[vertex])

        bron_kerbosch(
            R.union({vertex}),
            P.intersection(curr_neighbors),
            X.intersection(curr_neighbors),
            graph,
            max_clique
        )

        # Mark the current vertex as processed
        P.remove(vertex)
        X.add(vertex)


def main():

    if len(sys.argv) < 2:
        print("usage: python3 problem_2a.py filename")
        sys.exit(1)

    file_name = sys.argv[1]

    # Read in file, get graph, etc.
    graph, num_vertices = read_file_build_graph(file_name)

    start = time.time()

    # Initialize the max_clique set, R, P, and X.
    # Then, run bron_kerbosch to find the max clique.
    max_clique = [set()]
    R = set()
    P = set(range(num_vertices))
    X = set()
    bron_kerbosch(R, P, X, graph, max_clique)

    end = time.time()
    runtime = end - start

    print("Max clique: ", max_clique[0])
    print("Total time to find max clique: ", runtime, "sec")


if __name__ == "__main__":
    main()
