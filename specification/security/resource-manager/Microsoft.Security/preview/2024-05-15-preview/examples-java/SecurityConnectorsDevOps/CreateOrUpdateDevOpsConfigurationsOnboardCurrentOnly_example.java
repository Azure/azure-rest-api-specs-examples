
import com.azure.resourcemanager.security.fluent.models.DevOpsConfigurationInner;
import com.azure.resourcemanager.security.models.Authorization;
import com.azure.resourcemanager.security.models.AutoDiscovery;
import com.azure.resourcemanager.security.models.DevOpsConfigurationProperties;

/**
 * Samples for DevOpsConfigurations CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2024-05-15-preview/examples/
     * SecurityConnectorsDevOps/CreateOrUpdateDevOpsConfigurationsOnboardCurrentOnly_example.json
     */
    /**
     * Sample code: CreateOrUpdate_DevOpsConfigurations_OnboardCurrentOnly.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void createOrUpdateDevOpsConfigurationsOnboardCurrentOnly(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.devOpsConfigurations().createOrUpdate("myRg", "mySecurityConnectorName",
            new DevOpsConfigurationInner().withProperties(new DevOpsConfigurationProperties()
                .withAuthorization(new Authorization().withCode("fakeTokenPlaceholder"))
                .withAutoDiscovery(AutoDiscovery.DISABLED)),
            com.azure.core.util.Context.NONE);
    }
}
