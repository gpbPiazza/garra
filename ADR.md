# Architecture decision records


## 2.

---
## 1. We will couple extract data from the files using Keys markers.

### **1. Extracting Data Using X and Y Coordinates**
- **How It Works**: 
  - You analyze the PDF layout and determine the X and Y coordinates of the elements you want to extract.
  - Use a PDF library (e.g., `github.com/ledongthuc/pdf` or `unidoc/unipdf`) to access the text or elements at specific positions.

- **Advantages**:
  - Works well for highly structured and consistently formatted PDFs (e.g., invoices, forms, or reports).
  - Precise control over the extraction process.
  - Can handle cases where the data is not delimited by text markers (e.g., tables or graphical elements).

- **Disadvantages**:
  - Fragile: If the PDF layout changes (e.g., font size, margins, or alignment), the extraction logic may break.
  - Requires significant effort to analyze and map the coordinates for each type of document.
  - Not suitable for PDFs with dynamic or inconsistent layouts.

- **When to Use**:
  - Use this approach if the PDF documents are guaranteed to have a fixed layout and you need precise control over the extraction process.


### **2. Extracting Data Using StartKey and EndKey**
- **How It Works**:
  - You read the text content of the PDF and define "start keys" and "end keys" to delimit the data you want to extract.
  - Extract the text between these keys programmatically.

- **Advantages**:
  - More flexible: Works well for PDFs with dynamic or semi-structured layouts.
  - Easier to implement and maintain compared to the coordinate-based approach.
  - Resilient to minor layout changes (e.g., text alignment or spacing).

- **Disadvantages**:
  - Relies on the presence of consistent textual markers (start and end keys).
  - May require additional logic to handle edge cases (e.g., missing keys or overlapping data).

- **When to Use**:
  - Use this approach if the PDFs have consistent textual markers and you need a more robust solution that can handle minor layout changes.

### **Final Decision**
We will go by **Approach 2 (StartKey and EndKey)** as it is easier to implement, more flexible, and resilient to minor layout changes. If you encounter PDFs with highly structured layouts, you can combine it with **Approach 1 (X and Y Coordinates)** for specific cases where precision is required.

---