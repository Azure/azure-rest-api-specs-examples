/** Samples for ChildResources List. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/stable/2022-10-01/examples/ChildResources_List.json
     */
    /**
     * Sample code: GetCurrentChildHealthByResource.
     *
     * @param manager Entry point to ResourceHealthManager.
     */
    public static void getCurrentChildHealthByResource(
        com.azure.resourcemanager.resourcehealth.ResourceHealthManager manager) {
        manager.childResources().list("resourceUri", null, null, com.azure.core.util.Context.NONE);
    }
}
