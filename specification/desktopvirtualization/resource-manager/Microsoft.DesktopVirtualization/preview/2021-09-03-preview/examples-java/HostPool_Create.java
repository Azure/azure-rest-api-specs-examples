import com.azure.resourcemanager.desktopvirtualization.fluent.models.RegistrationInfoInner;
import com.azure.resourcemanager.desktopvirtualization.models.HostPoolType;
import com.azure.resourcemanager.desktopvirtualization.models.LoadBalancerType;
import com.azure.resourcemanager.desktopvirtualization.models.MigrationRequestProperties;
import com.azure.resourcemanager.desktopvirtualization.models.Operation;
import com.azure.resourcemanager.desktopvirtualization.models.PersonalDesktopAssignmentType;
import com.azure.resourcemanager.desktopvirtualization.models.PreferredAppGroupType;
import com.azure.resourcemanager.desktopvirtualization.models.PublicNetworkAccess;
import com.azure.resourcemanager.desktopvirtualization.models.RegistrationTokenOperation;
import com.azure.resourcemanager.desktopvirtualization.models.SsoSecretType;
import java.time.OffsetDateTime;
import java.util.HashMap;
import java.util.Map;

/** Samples for HostPools CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/HostPool_Create.json
     */
    /**
     * Sample code: HostPool_Create.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void hostPoolCreate(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager
            .hostPools()
            .define("hostPool1")
            .withRegion("centralus")
            .withExistingResourceGroup("resourceGroup1")
            .withHostPoolType(HostPoolType.POOLED)
            .withLoadBalancerType(LoadBalancerType.BREADTH_FIRST)
            .withPreferredAppGroupType(PreferredAppGroupType.DESKTOP)
            .withTags(mapOf("tag1", "value1", "tag2", "value2"))
            .withFriendlyName("friendly")
            .withDescription("des1")
            .withPersonalDesktopAssignmentType(PersonalDesktopAssignmentType.AUTOMATIC)
            .withMaxSessionLimit(999999)
            .withRegistrationInfo(
                new RegistrationInfoInner()
                    .withExpirationTime(OffsetDateTime.parse("2020-10-01T14:01:54.9571247Z"))
                    .withRegistrationTokenOperation(RegistrationTokenOperation.UPDATE))
            .withVmTemplate("{json:json}")
            .withSsoadfsAuthority("https://adfs")
            .withSsoClientId("client")
            .withSsoClientSecretKeyVaultPath("https://keyvault/secret")
            .withSsoSecretType(SsoSecretType.SHARED_KEY)
            .withStartVMOnConnect(false)
            .withMigrationRequest(
                new MigrationRequestProperties()
                    .withOperation(Operation.START)
                    .withMigrationPath(
                        "TenantGroups/{defaultV1TenantGroup.Name}/Tenants/{defaultV1Tenant.Name}/HostPools/{sessionHostPool.Name}"))
            .withPublicNetworkAccess(PublicNetworkAccess.ENABLED)
            .create();
    }

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
