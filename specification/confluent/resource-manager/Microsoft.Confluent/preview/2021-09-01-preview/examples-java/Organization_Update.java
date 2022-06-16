import com.azure.core.util.Context;
import com.azure.resourcemanager.confluent.models.OrganizationResource;
import java.util.HashMap;
import java.util.Map;

/** Samples for Organization Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/confluent/resource-manager/Microsoft.Confluent/preview/2021-09-01-preview/examples/Organization_Update.json
     */
    /**
     * Sample code: Confluent_Update.
     *
     * @param manager Entry point to ConfluentManager.
     */
    public static void confluentUpdate(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        OrganizationResource resource =
            manager
                .organizations()
                .getByResourceGroupWithResponse("myResourceGroup", "myOrganization", Context.NONE)
                .getValue();
        resource.update().withTags(mapOf("client", "dev-client", "env", "dev")).apply();
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
