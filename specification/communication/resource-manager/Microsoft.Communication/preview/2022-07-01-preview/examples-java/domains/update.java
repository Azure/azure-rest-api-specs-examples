import com.azure.core.util.Context;
import com.azure.resourcemanager.communication.models.DomainResource;
import com.azure.resourcemanager.communication.models.UserEngagementTracking;
import java.util.HashMap;
import java.util.Map;

/** Samples for Domains Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/preview/2022-07-01-preview/examples/domains/update.json
     */
    /**
     * Sample code: Update Domains resource.
     *
     * @param manager Entry point to CommunicationManager.
     */
    public static void updateDomainsResource(com.azure.resourcemanager.communication.CommunicationManager manager) {
        DomainResource resource =
            manager
                .domains()
                .getWithResponse("MyResourceGroup", "MyEmailServiceResource", "mydomain.com", Context.NONE)
                .getValue();
        resource
            .update()
            .withValidSenderUsernames(mapOf("info", "MyDomain Info", "alerts", "MyDomain Alerts"))
            .withUserEngagementTracking(UserEngagementTracking.ENABLED)
            .apply();
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
