import com.azure.resourcemanager.apimanagement.models.ApiManagementServiceResource;
import java.util.HashMap;
import java.util.Map;

/** Samples for ApiManagementService Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementUpdateServicePublisherDetails.json
     */
    /**
     * Sample code: ApiManagementUpdateServicePublisherDetails.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementUpdateServicePublisherDetails(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        ApiManagementServiceResource resource =
            manager
                .apiManagementServices()
                .getByResourceGroupWithResponse("rg1", "apimService1", com.azure.core.util.Context.NONE)
                .getValue();
        resource.update().withPublisherEmail("foobar@live.com").withPublisherName("Contoso Vnext").apply();
    }

    // Use "Map.of" if available
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
