import com.azure.resourcemanager.avs.models.PortMirroringDirectionEnum;

/** Samples for WorkloadNetworks CreatePortMirroring. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/WorkloadNetworks_CreatePortMirroringProfiles.json
     */
    /**
     * Sample code: WorkloadNetworks_CreatePortMirroring.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksCreatePortMirroring(com.azure.resourcemanager.avs.AvsManager manager) {
        manager
            .workloadNetworks()
            .definePortMirroring("portMirroring1")
            .withExistingPrivateCloud("group1", "cloud1")
            .withDisplayName("portMirroring1")
            .withDirection(PortMirroringDirectionEnum.BIDIRECTIONAL)
            .withSource("vmGroup1")
            .withDestination("vmGroup2")
            .withRevision(1L)
            .create();
    }
}
