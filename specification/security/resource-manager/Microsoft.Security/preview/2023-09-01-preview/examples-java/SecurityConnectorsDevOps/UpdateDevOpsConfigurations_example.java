
import com.azure.resourcemanager.security.fluent.models.DevOpsConfigurationInner;
import com.azure.resourcemanager.security.models.AutoDiscovery;
import com.azure.resourcemanager.security.models.DevOpsConfigurationProperties;

/**
 * Samples for DevOpsConfigurations Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2023-09-01-preview/examples/
     * SecurityConnectorsDevOps/UpdateDevOpsConfigurations_example.json
     */
    /**
     * Sample code: Update_DevOpsConfigurations.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void updateDevOpsConfigurations(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.devOpsConfigurations().update("myRg", "mySecurityConnectorName",
            new DevOpsConfigurationInner()
                .withProperties(new DevOpsConfigurationProperties().withAutoDiscovery(AutoDiscovery.ENABLED)),
            com.azure.core.util.Context.NONE);
    }
}
