
/**
 * Samples for SovereignViewOperations GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-31-preview/SovereignViewOperations_Get.json
     */
    /**
     * Sample code: SovereignViewOperations_Get.
     * 
     * @param manager Entry point to SovereignManager.
     */
    public static void sovereignViewOperationsGet(com.azure.resourcemanager.sovereign.SovereignManager manager) {
        manager.sovereignViewOperations().getByResourceGroupWithResponse("ExampleResourceGroup", "ExampleSV",
            com.azure.core.util.Context.NONE);
    }
}
