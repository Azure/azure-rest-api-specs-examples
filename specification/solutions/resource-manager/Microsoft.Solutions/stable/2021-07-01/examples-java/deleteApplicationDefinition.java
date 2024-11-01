
/**
 * Samples for ApplicationDefinitions DeleteById.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/
     * deleteApplicationDefinition.json
     */
    /**
     * Sample code: Deletes managed application definition.
     * 
     * @param manager Entry point to ApplicationManager.
     */
    public static void
        deletesManagedApplicationDefinition(com.azure.resourcemanager.managedapplications.ApplicationManager manager) {
        manager.applicationDefinitions().deleteByIdWithResponse("rg", "myManagedApplicationDef",
            com.azure.core.util.Context.NONE);
    }
}
