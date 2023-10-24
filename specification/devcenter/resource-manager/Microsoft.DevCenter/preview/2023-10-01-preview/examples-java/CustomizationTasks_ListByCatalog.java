/** Samples for CustomizationTasks ListByCatalog. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2023-10-01-preview/examples/CustomizationTasks_ListByCatalog.json
     */
    /**
     * Sample code: CustomizationTasks_ListByCatalog.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void customizationTasksListByCatalog(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager
            .customizationTasks()
            .listByCatalog("rg1", "Contoso", "CentralCatalog", null, com.azure.core.util.Context.NONE);
    }
}
