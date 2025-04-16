# Architecture decision records


## We will couple extract data from the files using Keys markers.

---

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

---

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

---

### **3. Alternative Approaches**
Here are some additional approaches you can consider:

#### **a. Using Regular Expressions**
- Extract text from the PDF and use regular expressions to identify and extract the desired data.
- **Advantages**: Works well for extracting patterns like dates, numbers, or specific formats (e.g., "Invoice #12345").
- **Disadvantages**: Requires well-defined patterns and may fail if the text format changes.

#### **b. Using Machine Learning (OCR or NLP)**
- Use Optical Character Recognition (OCR) tools (e.g., Tesseract) to extract text from scanned PDFs.
- Apply Natural Language Processing (NLP) techniques to identify and extract key information.
- **Advantages**: Works for unstructured or scanned PDFs without consistent formatting.
- **Disadvantages**: Complex to implement and may require training data for accurate results.

#### **c. Using PDF Libraries with Table Extraction**
- If the data is in tabular format, use libraries like `tabula` (Python) or `camelot` to extract tables directly.
- **Advantages**: Simplifies table extraction.
- **Disadvantages**: Limited to tabular data.

---

### **Recommendation**
- **If the PDFs are highly structured and consistent**: Use **Approach 1 (X and Y Coordinates)** for precise extraction.
- **If the PDFs have consistent textual markers**: Use **Approach 2 (StartKey and EndKey)** for a more flexible and robust solution.
- **If the PDFs are unstructured or scanned**: Consider **OCR/NLP** or **regular expressions** for extracting data.

---

### **Final Decision**
Start with **Approach 2 (StartKey and EndKey)** as it is easier to implement, more flexible, and resilient to minor layout changes. If you encounter PDFs with highly structured layouts, you can combine it with **Approach 1 (X and Y Coordinates)** for specific cases where precision is required.