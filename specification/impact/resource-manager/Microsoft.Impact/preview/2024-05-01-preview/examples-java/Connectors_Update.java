
import com.azure.resourcemanager.impactreporting.models.Connector;
import com.azure.resourcemanager.impactreporting.models.ConnectorUpdateProperties;
import com.azure.resourcemanager.impactreporting.models.Platform;

/**
 * Samples for Connectors Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-05-01-preview/Connectors_Update.json
     */
    /**
     * Sample code: Connectors_Update.
     * 
     * @param manager Entry point to ImpactReportingManager.
     */
    public static void connectorsUpdate(com.azure.resourcemanager.impactreporting.ImpactReportingManager manager) {
        Connector resource
            = manager.connectors().getWithResponse("testconnector1", com.azure.core.util.Context.NONE).getValue();
        resource.update().withProperties(new ConnectorUpdateProperties().withConnectorType(Platform.AZURE_MONITOR))
            .apply();
    }
}
