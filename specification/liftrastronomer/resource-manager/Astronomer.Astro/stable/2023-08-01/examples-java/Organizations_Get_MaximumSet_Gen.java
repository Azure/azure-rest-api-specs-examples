
/**
 * Samples for Organizations GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/liftrastronomer/resource-manager/Astronomer.Astro/stable/2023-08-01/examples/
     * Organizations_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: Organizations_Get - generated by [MaximumSet] rule.
     * 
     * @param manager Entry point to AstroManager.
     */
    public static void organizationsGetGeneratedByMaximumSetRule(com.azure.resourcemanager.astro.AstroManager manager) {
        manager.organizations().getByResourceGroupWithResponse("rgastronomer", "S PS",
            com.azure.core.util.Context.NONE);
    }
}
