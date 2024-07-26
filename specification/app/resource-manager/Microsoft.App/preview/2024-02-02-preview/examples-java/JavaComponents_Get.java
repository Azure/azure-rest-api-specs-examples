
/**
 * Samples for JavaComponents Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/JavaComponents_Get.json
     */
    /**
     * Sample code: Get Java Component.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void getJavaComponent(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.javaComponents().getWithResponse("examplerg", "myenvironment", "myjavacomponent",
            com.azure.core.util.Context.NONE);
    }
}
