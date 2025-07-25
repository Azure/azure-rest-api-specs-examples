
import com.azure.resourcemanager.dashboard.models.ApiKey;
import com.azure.resourcemanager.dashboard.models.AzureMonitorWorkspaceIntegration;
import com.azure.resourcemanager.dashboard.models.DeterministicOutboundIp;
import com.azure.resourcemanager.dashboard.models.EnterpriseConfigurations;
import com.azure.resourcemanager.dashboard.models.GrafanaConfigurations;
import com.azure.resourcemanager.dashboard.models.GrafanaIntegrations;
import com.azure.resourcemanager.dashboard.models.GrafanaPlugin;
import com.azure.resourcemanager.dashboard.models.ManagedGrafana;
import com.azure.resourcemanager.dashboard.models.ManagedGrafanaPropertiesUpdateParameters;
import com.azure.resourcemanager.dashboard.models.MarketplaceAutoRenew;
import com.azure.resourcemanager.dashboard.models.ResourceSku;
import com.azure.resourcemanager.dashboard.models.Security;
import com.azure.resourcemanager.dashboard.models.Smtp;
import com.azure.resourcemanager.dashboard.models.Snapshots;
import com.azure.resourcemanager.dashboard.models.StartTlsPolicy;
import com.azure.resourcemanager.dashboard.models.UnifiedAlertingScreenshots;
import com.azure.resourcemanager.dashboard.models.Users;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Grafana Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01-preview/Grafana_Update.json
     */
    /**
     * Sample code: Grafana_Update.
     * 
     * @param manager Entry point to DashboardManager.
     */
    public static void grafanaUpdate(com.azure.resourcemanager.dashboard.DashboardManager manager) {
        ManagedGrafana resource = manager.grafanas()
            .getByResourceGroupWithResponse("myResourceGroup", "myWorkspace", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withTags(mapOf("Environment", "Dev 2")).withSku(new ResourceSku().withName("Standard"))
            .withProperties(new ManagedGrafanaPropertiesUpdateParameters().withApiKey(ApiKey.ENABLED)
                .withDeterministicOutboundIp(DeterministicOutboundIp.ENABLED)
                .withGrafanaIntegrations(new GrafanaIntegrations().withAzureMonitorWorkspaceIntegrations(
                    Arrays.asList(new AzureMonitorWorkspaceIntegration().withAzureMonitorWorkspaceResourceId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.monitor/accounts/myAzureMonitorWorkspace"))))
                .withEnterpriseConfigurations(new EnterpriseConfigurations().withMarketplacePlanId("myPlanId")
                    .withMarketplaceAutoRenew(MarketplaceAutoRenew.ENABLED))
                .withGrafanaConfigurations(new GrafanaConfigurations()
                    .withSmtp(new Smtp().withEnabled(true).withHost("smtp.sendemail.com:587").withUser("username")
                        .withPassword("fakeTokenPlaceholder").withFromAddress("test@sendemail.com")
                        .withFromName("emailsender").withStartTlsPolicy(StartTlsPolicy.OPPORTUNISTIC_START_TLS)
                        .withSkipVerify(true))
                    .withSnapshots(new Snapshots().withExternalEnabled(true))
                    .withUsers(new Users().withViewersCanEdit(true).withEditorsCanAdmin(true))
                    .withSecurity(new Security().withCsrfAlwaysCheck(false))
                    .withUnifiedAlertingScreenshots(new UnifiedAlertingScreenshots().withCaptureEnabled(false)))
                .withGrafanaPlugins(mapOf("sample-plugin-id", new GrafanaPlugin())).withGrafanaMajorVersion("9"))
            .apply();
    }

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
