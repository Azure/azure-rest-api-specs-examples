import com.azure.resourcemanager.automation.models.FieldDefinition;
import java.util.HashMap;
import java.util.Map;

/** Samples for ConnectionType CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/createOrUpdateConnectionType.json
     */
    /**
     * Sample code: Create or update connection type.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void createOrUpdateConnectionType(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .connectionTypes()
            .define("myCT")
            .withExistingAutomationAccount("rg", "myAutomationAccount22")
            .withName("myCT")
            .withFieldDefinitions(
                mapOf(
                    "myBoolField",
                    new FieldDefinition().withIsEncrypted(false).withIsOptional(false).withType("bool"),
                    "myStringField",
                    new FieldDefinition().withIsEncrypted(false).withIsOptional(false).withType("string"),
                    "myStringFieldEncrypted",
                    new FieldDefinition().withIsEncrypted(true).withIsOptional(false).withType("string")))
            .withIsGlobal(false)
            .create();
    }

    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
