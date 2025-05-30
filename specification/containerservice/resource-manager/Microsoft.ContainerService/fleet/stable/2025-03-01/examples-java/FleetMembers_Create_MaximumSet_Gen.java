
/**
 * Samples for FleetMembers Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/FleetMembers_Create_MaximumSet_Gen.json
     */
    /**
     * Sample code: Creates a FleetMember resource with a long running operation. - generated by [MaximumSet] rule.
     * 
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void createsAFleetMemberResourceWithALongRunningOperationGeneratedByMaximumSetRule(
        com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.fleetMembers().define("fleet1").withExistingFleet("rgfleets", "fleet1").withClusterResourceId(
            "/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/cluster-1")
            .withGroup("fleet1").withIfMatch("amkttadbw").withIfNoneMatch("zoljoccbcg").create();
    }
}
