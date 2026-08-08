
/**
 * Samples for AIManagers GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-02-preview/AIManagers_Get.json
     */
    /**
     * Sample code: Gets an AI Manager resource.
     * 
     * @param manager Entry point to ContainerServiceAIManagerManager.
     */
    public static void getsAnAIManagerResource(
        com.azure.resourcemanager.containerserviceaimanager.ContainerServiceAIManagerManager manager) {
        manager.aIManagers().getByResourceGroupWithResponse("rg1", "aimanager1", com.azure.core.util.Context.NONE);
    }
}
