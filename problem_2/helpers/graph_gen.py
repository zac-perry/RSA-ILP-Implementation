import networkx as nx
import random


def generate_random_graph_file(n_vertices, edge_probability, output_file):
    # Generate a random Erdős-Rényi graph with n vertices
    G = nx.erdos_renyi_graph(n_vertices, edge_probability)

    # Count the number of edges
    n_edges = G.number_of_edges()

    # Write to file in the modified DIMACS format
    with open(output_file, "w") as f:
        f.write(f"{n_vertices} {n_edges}\n")

        # Write each edge with a random weight
        for u, v in G.edges():
            f.write(f"{u} {v}\n")


# Generate a random graph with 1000 vertices and a medium density
# The edge_probability controls how dense the graph is (0.01 = ~1% of possible edges exist)
generate_random_graph_file(1000, 0.01, "input/random_graph_1000.txt")

print("Random graph file created successfully!")
