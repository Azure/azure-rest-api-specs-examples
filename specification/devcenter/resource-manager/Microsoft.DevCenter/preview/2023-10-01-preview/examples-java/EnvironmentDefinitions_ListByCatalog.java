/** Samples for EnvironmentDefinitions ListByCatalog. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2023-10-01-preview/examples/EnvironmentDefinitions_ListByCatalog.json
     */
    /**
     * Sample code: EnvironmentDefinitions_ListByCatalog.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void environmentDefinitionsListByCatalog(
        com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager
            .environmentDefinitions()
            .listByCatalog("rg1", "Contoso", "myCatalog", null, com.azure.core.util.Context.NONE);
    }
}
