
import com.azure.resourcemanager.deviceregistry.models.SchemaVersionProperties;

/**
 * Samples for SchemaVersions CreateOrReplace.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01-preview/Create_SchemaVersion.json
     */
    /**
     * Sample code: Create_SchemaVersion.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void createSchemaVersion(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.schemaVersions().define("1").withExistingSchema("myResourceGroup", "my-schema-registry", "my-schema")
            .withProperties(new SchemaVersionProperties().withDescription("Schema version 1").withSchemaContent(
                "{\"$schema\": \"http://json-schema.org/draft-07/schema#\",\"type\": \"object\",\"properties\": {\"humidity\": {\"type\": \"string\"},\"temperature\": {\"type\":\"number\"}}}"))
            .create();
    }
}
