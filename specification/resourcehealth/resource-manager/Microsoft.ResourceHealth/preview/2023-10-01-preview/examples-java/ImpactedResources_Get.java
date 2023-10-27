/** Samples for ImpactedResources Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/preview/2023-10-01-preview/examples/ImpactedResources_Get.json
     */
    /**
     * Sample code: ImpactedResourcesGet.
     *
     * @param manager Entry point to ResourceHealthManager.
     */
    public static void impactedResourcesGet(com.azure.resourcemanager.resourcehealth.ResourceHealthManager manager) {
        manager.impactedResources().getWithResponse("BC_1-FXZ", "abc-123-ghj-456", com.azure.core.util.Context.NONE);
    }
}
