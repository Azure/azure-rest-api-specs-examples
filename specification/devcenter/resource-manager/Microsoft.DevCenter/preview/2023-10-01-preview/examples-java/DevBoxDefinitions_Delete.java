/** Samples for DevBoxDefinitions Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2023-10-01-preview/examples/DevBoxDefinitions_Delete.json
     */
    /**
     * Sample code: DevBoxDefinitions_Delete.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void devBoxDefinitionsDelete(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.devBoxDefinitions().delete("rg1", "Contoso", "WebDevBox", com.azure.core.util.Context.NONE);
    }
}
