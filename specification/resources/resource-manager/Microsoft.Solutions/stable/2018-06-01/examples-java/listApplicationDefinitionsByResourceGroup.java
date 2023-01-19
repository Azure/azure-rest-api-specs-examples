/** Samples for ApplicationDefinitions ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Solutions/stable/2018-06-01/examples/listApplicationDefinitionsByResourceGroup.json
     */
    /**
     * Sample code: List managed application definitions.
     *
     * @param manager Entry point to ApplicationManager.
     */
    public static void listManagedApplicationDefinitions(
        com.azure.resourcemanager.managedapplications.ApplicationManager manager) {
        manager.applicationDefinitions().listByResourceGroup("rg", com.azure.core.util.Context.NONE);
    }
}
