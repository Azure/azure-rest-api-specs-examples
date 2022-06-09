```java
import com.azure.resourcemanager.apimanagement.models.SchemaType;

/** Samples for GlobalSchema CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateGlobalSchema1.json
     */
    /**
     * Sample code: ApiManagementCreateSchema1.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateSchema1(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .globalSchemas()
            .define("schema1")
            .withExistingService("rg1", "apimService1")
            .withSchemaType(SchemaType.XML)
            .withDescription("sample schema description")
            .withValue(
                "<xsd:schema xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\"\r\n"
                    + "           xmlns:tns=\"http://tempuri.org/PurchaseOrderSchema.xsd\"\r\n"
                    + "           targetNamespace=\"http://tempuri.org/PurchaseOrderSchema.xsd\"\r\n"
                    + "           elementFormDefault=\"qualified\">\r\n"
                    + " <xsd:element name=\"PurchaseOrder\" type=\"tns:PurchaseOrderType\"/>\r\n"
                    + " <xsd:complexType name=\"PurchaseOrderType\">\r\n"
                    + "  <xsd:sequence>\r\n"
                    + "   <xsd:element name=\"ShipTo\" type=\"tns:USAddress\" maxOccurs=\"2\"/>\r\n"
                    + "   <xsd:element name=\"BillTo\" type=\"tns:USAddress\"/>\r\n"
                    + "  </xsd:sequence>\r\n"
                    + "  <xsd:attribute name=\"OrderDate\" type=\"xsd:date\"/>\r\n"
                    + " </xsd:complexType>\r\n\r\n"
                    + " <xsd:complexType name=\"USAddress\">\r\n"
                    + "  <xsd:sequence>\r\n"
                    + "   <xsd:element name=\"name\"   type=\"xsd:string\"/>\r\n"
                    + "   <xsd:element name=\"street\" type=\"xsd:string\"/>\r\n"
                    + "   <xsd:element name=\"city\"   type=\"xsd:string\"/>\r\n"
                    + "   <xsd:element name=\"state\"  type=\"xsd:string\"/>\r\n"
                    + "   <xsd:element name=\"zip\"    type=\"xsd:integer\"/>\r\n"
                    + "  </xsd:sequence>\r\n"
                    + "  <xsd:attribute name=\"country\" type=\"xsd:NMTOKEN\" fixed=\"US\"/>\r\n"
                    + " </xsd:complexType>\r\n"
                    + "</xsd:schema>")
            .create();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.
