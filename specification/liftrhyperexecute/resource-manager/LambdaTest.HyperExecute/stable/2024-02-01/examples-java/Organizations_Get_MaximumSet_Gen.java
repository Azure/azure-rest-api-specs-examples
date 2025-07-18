
/**
 * Samples for Organizations GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-02-01/Organizations_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: Organizations_Get_MaximumSet_Gen - generated by [MaximumSet] rule - generated by [MaximumSet] rule -
     * generated by [MaximumSet] rule.
     * 
     * @param manager Entry point to LambdaTestHyperExecuteManager.
     */
    public static void
        organizationsGetMaximumSetGenGeneratedByMaximumSetRuleGeneratedByMaximumSetRuleGeneratedByMaximumSetRule(
            com.azure.resourcemanager.lambdatesthyperexecute.LambdaTestHyperExecuteManager manager) {
        manager.organizations().getByResourceGroupWithResponse("rgopenapi", "testorg",
            com.azure.core.util.Context.NONE);
    }
}
