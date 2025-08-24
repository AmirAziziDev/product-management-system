## Objective
- Develop a REST API that reads information about products from database. You can use C#,
  JavaScript or GO. You can use any database engine.
- Develop a frontend for it using React or any other library you prefer.
- You can use any libraries or tools you prefer.
- Publish your project to Github and be ready to share it. Preferably keep git history of your work.
- Share link to Github one working day before interview to recruiter.
## Data model
The database should be capable of storing information about the following entities:
- *Product*
  Each product should have a Name (text), a Product Type, and multiple associated Colours.
- *Product Type*
- *Colour*
  For example, a Product with id = 1 could have Product Type = Sofa and Colours = Blue, Ruby.
  You may add additional tables and columns as needed.
## Sample Data Loading
Load the contents of product-types.txt into the ProductType table. Do the same for colours.txt
and Colour table.
Internal
## API Endpoints
Create the following endpoints:
- Add a new product. The payload must contain: name, product type id, colour id.
- Returns a list of products. The payload must contain: id, name of product.
- Returns details about a product by id. The payload must contain: id, name, product type,
  colours associated with that product.
## User interface
The user interface should have two pages:
- Create a New Product: A page to create a new product.
- View Products List: A page displaying a list of products in a table with the following columns:
  Product Id, Product Name, Product Type, and a comma-separated list of Colours. The data
  should be ordered by the creation date, with the newest entries at the top.
  Internal