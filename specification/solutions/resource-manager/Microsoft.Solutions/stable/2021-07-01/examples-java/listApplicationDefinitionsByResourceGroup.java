/** Samples for ApplicationDefinitions ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/listApplicationDefinitionsByResourceGroup.json
     */
    /**
     * Sample code: Lists the managed application definitions in a resource group.
     *
     * @param manager Entry point to ApplicationManager.
     */
    public static void listsTheManagedApplicationDefinitionsInAResourceGroup(
        com.azure.resourcemanager.managedapplications.ApplicationManager manager) {
        manager.applicationDefinitions().listByResourceGroup("rg", com.azure.core.util.Context.NONE);
    }
}
