
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import java.io.IOException;

/**
 * Samples for ContentType CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementCreateContentType.json
     */
    /**
     * Sample code: ApiManagementCreateContentType.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateContentType(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) throws IOException {
        manager.contentTypes().define("page").withExistingService("rg1", "apimService1").withNamePropertiesName("Page")
            .withDescription("A regular page")
            .withSchema(SerializerFactory.createDefaultManagementSerializerAdapter().deserialize(
                "{\"additionalProperties\":false,\"properties\":{\"en_us\":{\"type\":\"object\",\"additionalProperties\":false,\"properties\":{\"description\":{\"type\":\"string\",\"description\":\"Page description. This property gets included in SEO attributes.\",\"indexed\":true,\"title\":\"Description\"},\"documentId\":{\"type\":\"string\",\"description\":\"Reference to page content document.\",\"title\":\"Document ID\"},\"keywords\":{\"type\":\"string\",\"description\":\"Page keywords. This property gets included in SEO attributes.\",\"indexed\":true,\"title\":\"Keywords\"},\"permalink\":{\"type\":\"string\",\"description\":\"Page permalink, e.g. '/about'.\",\"indexed\":true,\"title\":\"Permalink\"},\"title\":{\"type\":\"string\",\"description\":\"Page title. This property gets included in SEO attributes.\",\"indexed\":true,\"title\":\"Title\"}},\"required\":[\"title\",\"permalink\",\"documentId\"]}}}",
                Object.class, SerializerEncoding.JSON))
            .withVersion("1.0.0").create();
    }
}
