
/**
 * Samples for ParameterGroups List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/ParameterGroups_ListBySubscription.json
     */
    /**
     * Sample code: List HorizonDB parameter groups in a subscription.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void
        listHorizonDBParameterGroupsInASubscription(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.parameterGroups().list(com.azure.core.util.Context.NONE);
    }
}
