
/**
 * Samples for AppServiceEnvironments GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/AppServiceEnvironments_Get.json
     */
    /**
     * Sample code: Get the properties of an App Service Environment.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        getThePropertiesOfAnAppServiceEnvironment(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getAppServiceEnvironments().getByResourceGroupWithResponse("test-rg", "test-ase",
            com.azure.core.util.Context.NONE);
    }
}
