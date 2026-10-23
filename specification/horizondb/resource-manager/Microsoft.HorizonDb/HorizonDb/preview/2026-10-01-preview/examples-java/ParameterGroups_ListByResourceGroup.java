
/**
 * Samples for ParameterGroups ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/ParameterGroups_ListByResourceGroup.json
     */
    /**
     * Sample code: List HorizonDB parameter groups in a resource group.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void
        listHorizonDBParameterGroupsInAResourceGroup(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.parameterGroups().listByResourceGroup("example-resource-group", com.azure.core.util.Context.NONE);
    }
}
