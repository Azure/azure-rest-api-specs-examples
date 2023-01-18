/** Samples for ApplicationDefinitions GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Solutions/stable/2018-06-01/examples/getApplicationDefinition.json
     */
    /**
     * Sample code: Get managed application definition.
     *
     * @param manager Entry point to ApplicationManager.
     */
    public static void getManagedApplicationDefinition(
        com.azure.resourcemanager.managedapplications.ApplicationManager manager) {
        manager
            .applicationDefinitions()
            .getByResourceGroupWithResponse("rg", "myManagedApplicationDef", com.azure.core.util.Context.NONE);
    }
}
