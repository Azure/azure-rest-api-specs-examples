
/**
 * Samples for Monitors ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01-preview/Monitors_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: Monitors_ListByResourceGroup_MaximumSet_Gen.
     * 
     * @param manager Entry point to NewRelicObservabilityManager.
     */
    public static void monitorsListByResourceGroupMaximumSetGen(
        com.azure.resourcemanager.newrelicobservability.NewRelicObservabilityManager manager) {
        manager.monitors().listByResourceGroup("rgNewRelic", com.azure.core.util.Context.NONE);
    }
}
