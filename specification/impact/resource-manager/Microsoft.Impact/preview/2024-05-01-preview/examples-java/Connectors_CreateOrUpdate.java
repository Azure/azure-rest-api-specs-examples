
import com.azure.resourcemanager.impactreporting.models.ConnectorProperties;
import com.azure.resourcemanager.impactreporting.models.Platform;

/**
 * Samples for Connectors CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-05-01-preview/Connectors_CreateOrUpdate.json
     */
    /**
     * Sample code: Connectors_CreateOrUpdate.
     * 
     * @param manager Entry point to ImpactReportingManager.
     */
    public static void
        connectorsCreateOrUpdate(com.azure.resourcemanager.impactreporting.ImpactReportingManager manager) {
        manager.connectors().define("testconnector1")
            .withProperties(new ConnectorProperties().withConnectorType(Platform.AZURE_MONITOR)).create();
    }
}
