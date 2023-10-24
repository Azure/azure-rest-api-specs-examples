/** Samples for CustomizationTasks Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2023-10-01-preview/examples/CustomizationTasks_Get.json
     */
    /**
     * Sample code: CustomizationTasks_Get.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void customizationTasksGet(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager
            .customizationTasks()
            .getWithResponse("rg1", "Contoso", "CentralCatalog", "SampleTask", com.azure.core.util.Context.NONE);
    }
}
