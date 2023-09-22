/** Samples for ApplicationDefinitions GetById. */
public final class Main {
    /*
     * x-ms-original-file: specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/getApplicationDefinition.json
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
            .getByIdWithResponse("rg", "myManagedApplicationDef", com.azure.core.util.Context.NONE);
    }
}
