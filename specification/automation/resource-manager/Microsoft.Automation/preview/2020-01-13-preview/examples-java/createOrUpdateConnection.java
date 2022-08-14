import com.azure.resourcemanager.automation.models.ConnectionTypeAssociationProperty;
import java.util.HashMap;
import java.util.Map;

/** Samples for Connection CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/createOrUpdateConnection.json
     */
    /**
     * Sample code: Create or update connection.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void createOrUpdateConnection(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .connections()
            .define("mysConnection")
            .withExistingAutomationAccount("rg", "myAutomationAccount28")
            .withName("mysConnection")
            .withConnectionType(new ConnectionTypeAssociationProperty().withName("Azure"))
            .withDescription("my description goes here")
            .withFieldDefinitionValues(
                mapOf("AutomationCertificateName", "mysCertificateName", "SubscriptionID", "subid"))
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
