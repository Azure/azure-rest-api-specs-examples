
/**
 * Samples for DeidServices ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/healthdataaiservices/HealthDataAIServices.Management/examples/2024-02-28-preview/
     * DeidServices_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: DeidServices_ListByResourceGroup - generated by [MaximumSet] rule - generated by [MaximumSet] rule.
     * 
     * @param manager Entry point to HealthDataAIServicesManager.
     */
    public static void deidServicesListByResourceGroupGeneratedByMaximumSetRuleGeneratedByMaximumSetRule(
        com.azure.resourcemanager.healthdataaiservices.HealthDataAIServicesManager manager) {
        manager.deidServices().listByResourceGroup("rgopenapi", com.azure.core.util.Context.NONE);
    }
}
