
import com.azure.resourcemanager.security.fluent.models.DevOpsConfigurationInner;
import com.azure.resourcemanager.security.models.Authorization;
import com.azure.resourcemanager.security.models.AutoDiscovery;
import com.azure.resourcemanager.security.models.DevOpsConfigurationProperties;
import java.util.Arrays;

/**
 * Samples for DevOpsConfigurations CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2024-05-15-preview/examples/
     * SecurityConnectorsDevOps/CreateOrUpdateDevOpsConfigurationsOnboardSelected_example.json
     */
    /**
     * Sample code: CreateOrUpdate_DevOpsConfigurations_OnboardSelected.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void
        createOrUpdateDevOpsConfigurationsOnboardSelected(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.devOpsConfigurations().createOrUpdate("myRg", "mySecurityConnectorName",
            new DevOpsConfigurationInner().withProperties(new DevOpsConfigurationProperties()
                .withAuthorization(new Authorization().withCode("fakeTokenPlaceholder"))
                .withAutoDiscovery(AutoDiscovery.DISABLED).withTopLevelInventoryList(Arrays.asList("org1", "org2"))),
            com.azure.core.util.Context.NONE);
    }
}
