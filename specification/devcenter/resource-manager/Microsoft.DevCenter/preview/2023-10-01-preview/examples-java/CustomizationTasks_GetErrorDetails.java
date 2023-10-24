/** Samples for CustomizationTasks GetErrorDetails. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2023-10-01-preview/examples/CustomizationTasks_GetErrorDetails.json
     */
    /**
     * Sample code: CustomizationTasks_GetErrorDetails.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void customizationTasksGetErrorDetails(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager
            .customizationTasks()
            .getErrorDetailsWithResponse(
                "rg1", "Contoso", "CentralCatalog", "SampleTask", com.azure.core.util.Context.NONE);
    }
}
