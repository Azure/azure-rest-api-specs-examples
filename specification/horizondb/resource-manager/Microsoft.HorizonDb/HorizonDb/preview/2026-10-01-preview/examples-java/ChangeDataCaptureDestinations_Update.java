
import com.azure.resourcemanager.horizondb.models.ChangeDataCaptureDesiredState;
import com.azure.resourcemanager.horizondb.models.ChangeDataCaptureDestination;
import com.azure.resourcemanager.horizondb.models.ChangeDataCaptureDestinationPropertiesForPatchUpdate;

/**
 * Samples for ChangeDataCaptureDestinations Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/ChangeDataCaptureDestinations_Update.json
     */
    /**
     * Sample code: Disable a Microsoft Fabric change data capture destination.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void disableAMicrosoftFabricChangeDataCaptureDestination(
        com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        ChangeDataCaptureDestination resource
            = manager.changeDataCaptureDestinations().getWithResponse("example-resource-group", "example-cluster",
                "fabric-mirroring", com.azure.core.util.Context.NONE).getValue();
        resource.update().withProperties(new ChangeDataCaptureDestinationPropertiesForPatchUpdate()
            .withDesiredState(ChangeDataCaptureDesiredState.DISABLED)).apply();
    }
}
