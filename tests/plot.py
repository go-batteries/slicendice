import pandas as pd
import matplotlib.pyplot as plt

# Load data from the extracted files
data_half = pd.read_csv("data_half.txt", sep="\t")
data_full = pd.read_csv("data_full.txt", sep="\t")

data_half["Instruction/Cycle"] = data_half["Instructions"] / data_half["Cycles"]
data_full["Instruction/Cycle"] = data_full["Instructions"] / data_full["Cycles"]

data_half.columns = data_half.columns.str.strip()
data_full.columns = data_full.columns.str.strip()


data_half = data_half.drop(columns=["Benchmark"])
data_full = data_full.drop(columns=["Benchmark"])
# data = pd.concat([data_half, data_full], ignore_index=True)

# Convert columns to numeric where applicable
# Display column names to check for leading/trailing whitespace
print("\nData Half Columns:")
print(data_half, data_half.columns, data_half.values)
print("\nData Full Columns:")
print(data_full)  # Combine data into a single DataFrame for easier comparison

# Plotting
# # Plot cache references and misses
# plt.subplot(2, 2, 1)
# half_refs = data[data["Type"] == "ReverseHalf"]["CacheReferences"]
# half_misses = data[data["Type"] == "ReverseHalf"]["CacheMisses"]
# full_refs = data[data["Type"] == "ReverseFull"]["CacheReferences"]
# full_misses = data[data["Type"] == "ReverseFull"]["CacheMisses"]
# bar_width = 0.35
# index = range(len(half_refs))

# print("half refs")
# print(half_refs)
# print("half missess")
# print(half_misses)
# print("full missess")
# print(full_misses)
# print("full refs")
# print(full_refs)

plotframe = pd.DataFrame(
    {
        "Half": data_half.iloc[0].tolist(),
        "Full": data_full.iloc[0].tolist(),
    }
)

plt.figure(figsize=(14, 8))

labels = [
    "CacheReferences",
    "CacheMisses",
    "L1DcacheLoads",
    "L1DcacheMisses",
    "LLCLoads",
    "LLCMisses",
    "Branches",
    "Instruction/Cycle",
]

colors = [
    "r",
    "g",
    "b",
    "y",
    "m",
    "k",
    "orange",
    "purple",
]

reversed = list(reversed(colors))

# Plot data from 'data_half' (for "ReverseHalf")
for i, column in enumerate(labels):
    if column in data_half.columns:
        plt.plot(
            data_half.index,
            data_half[column],
            marker="o",
            label=f"{column} (Half)",
            color=colors[i],
            linestyle="-",
            alpha=0.7,
            markersize=20,
        )

# Plot data from 'data_full' (for "ReverseFull")
for i, column in enumerate(labels):
    if column in data_full.columns:
        plt.plot(
            data_full.index,
            data_full[column],
            marker="x",
            label=f"{column} (Full)",
            color=reversed[i],
            linestyle="-",
            alpha=0.7,
            markersize=15,
        )

# Add labels and title
plt.xlabel("Benchmark Instance")
plt.ylabel("Values")
plt.title("Comparison of ReverseHalf vs ReverseFull Data")

# Rotate x-axis labels for better readability
plt.xticks(rotation=45)

# Display the legend
plt.legend()

# Display grid for clarity
plt.grid(True)

# Show the plot
plt.tight_layout()
plt.show()

# pltdata_half = {
#     "CacheReferences": {"marker": "o", "label": "Cache References (Half)"},
#     "CacheMisses": {"marker": "o", "label": "Cache Misses (Half)"},
#     "L1DcacheLoads": {"marker": "o", "label": "L1Loads (Half)"},
#     "L1DcacheMisses": {"marker": "o", "label": "L1Misses (Half)"},
#     "LLCLoads": {"marker": "o", "label": "LLCLoads (Half)"},
#     "LLCMisses": {"marker": "o", "label": "LLCMisses (Half)"},
#     "Instruction/Cycle": {"marker": "o", "label": "IPC (Half)"},
# }
#
# pltdata_full = {
#     "CacheReferences": {"marker": "x", "label": "Cache References (Full)"},
#     "CacheMisses": {"marker": "x", "label": "Cache Misses (Full)"},
#     "L1DcacheLoads": {"marker": "x", "label": "L1Loads (Full)"},
#     "L1DcacheMisses": {"marker": "x", "label": "L1Misses (Full)"},
#     "LLCLoads": {"marker": "x", "label": "LLCLoads (Full)"},
#     "LLCMisses": {"marker": "x", "label": "LLCMisses (Full)"},
#     "Instruction/Cycle": {"marker": "x", "label": "IPC (Full)"},
# }
#
#
# # Axis labels and legend
# plt.xlabel("Benchmark Instance")
# plt.ylabel("Counts")
# plt.title("Cache References and Misses for ReverseHalf and ReverseFull")
# plt.legend()
# # plt.grid(visible=True)
# plt.plot(plotframe["Half"], plotframe["Full"])
# # Show plot
# plt.tight_layout()
# plt.show()
