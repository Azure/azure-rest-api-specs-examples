
import com.azure.resourcemanager.horizondb.models.ChangeDataCaptureDesiredState;
import com.azure.resourcemanager.horizondb.models.ChangeDataCaptureDestinationProperties;
import com.azure.resourcemanager.horizondb.models.FabricDestinationDetails;

/**
 * Samples for ChangeDataCaptureDestinations CreateOrReplace.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/ChangeDataCaptureDestinations_CreateOrReplace.json
     */
    /**
     * Sample code: Create a Microsoft Fabric change data capture destination.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void createAMicrosoftFabricChangeDataCaptureDestination(
        com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.changeDataCaptureDestinations().define("fabric-mirroring")
            .withExistingCluster("example-resource-group", "example-cluster")
            .withProperties(new ChangeDataCaptureDestinationProperties()
                .withDestination(new FabricDestinationDetails().withWorkspaceId("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"))
                .withDesiredState(ChangeDataCaptureDesiredState.ENABLED))
            .create();
    }
}
