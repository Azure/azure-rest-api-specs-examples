import com.azure.resourcemanager.scvmm.models.ExtendedLocation;

/** Samples for AvailabilitySets CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/CreateAvailabilitySet.json
     */
    /**
     * Sample code: CreateAvailabilitySet.
     *
     * @param manager Entry point to ScvmmManager.
     */
    public static void createAvailabilitySet(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager
            .availabilitySets()
            .define("HRAvailabilitySet")
            .withRegion("East US")
            .withExistingResourceGroup("testrg")
            .withExtendedLocation(
                new ExtendedLocation()
                    .withType("customLocation")
                    .withName(
                        "/subscriptions/a5015e1c-867f-4533-8541-85cd470d0cfb/resourceGroups/demoRG/providers/Microsoft.Arc/customLocations/contoso"))
            .withAvailabilitySetName("hr-avset")
            .withVmmServerId(
                "/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ScVmm/VMMServers/ContosoVMMServer")
            .create();
    }
}
