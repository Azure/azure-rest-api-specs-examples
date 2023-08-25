import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.apimanagement.models.SchemaType;
import java.io.IOException;

/** Samples for GlobalSchema CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateGlobalSchema2.json
     */
    /**
     * Sample code: ApiManagementCreateSchema2.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateSchema2(com.azure.resourcemanager.apimanagement.ApiManagementManager manager)
        throws IOException {
        manager
            .globalSchemas()
            .define("schema1")
            .withExistingService("rg1", "apimService1")
            .withSchemaType(SchemaType.JSON)
            .withDescription("sample schema description")
            .withDocument(
                SerializerFactory
                    .createDefaultManagementSerializerAdapter()
                    .deserialize(
                        "{\"type\":\"object\",\"$id\":\"https://example.com/person.schema.json\",\"$schema\":\"https://json-schema.org/draft/2020-12/schema\",\"properties\":{\"age\":{\"type\":\"integer\",\"description\":\"Age"
                            + " in years which must be equal to or greater than"
                            + " zero.\",\"minimum\":0},\"firstName\":{\"type\":\"string\",\"description\":\"The"
                            + " person's first name.\"},\"lastName\":{\"type\":\"string\",\"description\":\"The"
                            + " person's last name.\"}},\"title\":\"Person\"}",
                        Object.class,
                        SerializerEncoding.JSON))
            .create();
    }
}
