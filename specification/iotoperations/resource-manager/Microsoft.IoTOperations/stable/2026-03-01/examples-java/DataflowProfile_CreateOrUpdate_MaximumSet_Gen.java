
import com.azure.resourcemanager.iotoperations.models.DataflowProfileProperties;
import com.azure.resourcemanager.iotoperations.models.DiagnosticsLogs;
import com.azure.resourcemanager.iotoperations.models.ExtendedLocation;
import com.azure.resourcemanager.iotoperations.models.ExtendedLocationType;
import com.azure.resourcemanager.iotoperations.models.Metrics;
import com.azure.resourcemanager.iotoperations.models.ProfileDiagnostics;

/**
 * Samples for DataflowProfile CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/DataflowProfile_CreateOrUpdate_MaximumSet_Gen.json
     */
    /**
     * Sample code: DataflowProfile_CreateOrUpdate.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void
        dataflowProfileCreateOrUpdate(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.dataflowProfiles().define("resource-name123")
            .withExistingInstance("rgiotoperations", "resource-name123")
            .withProperties(new DataflowProfileProperties().withDiagnostics(
                new ProfileDiagnostics().withLogs(new DiagnosticsLogs().withLevel("rnmwokumdmebpmfxxxzvvjfdywotav"))
                    .withMetrics(new Metrics().withPrometheusPort(7581)))
                .withInstanceCount(14))
            .withExtendedLocation(new ExtendedLocation().withName(
                "/subscriptions/F8C729F9-DF9C-4743-848F-96EE433D8E53/resourceGroups/rgiotoperations/providers/Microsoft.ExtendedLocation/customLocations/resource-123")
                .withType(ExtendedLocationType.CUSTOM_LOCATION))
            .create();
    }
}
