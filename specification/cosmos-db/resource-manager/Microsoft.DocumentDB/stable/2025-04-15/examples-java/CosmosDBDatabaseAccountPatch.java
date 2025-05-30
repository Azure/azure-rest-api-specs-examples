
import com.azure.resourcemanager.cosmos.models.AnalyticalStorageConfiguration;
import com.azure.resourcemanager.cosmos.models.AnalyticalStorageSchemaType;
import com.azure.resourcemanager.cosmos.models.BackupStorageRedundancy;
import com.azure.resourcemanager.cosmos.models.Capacity;
import com.azure.resourcemanager.cosmos.models.ConsistencyPolicy;
import com.azure.resourcemanager.cosmos.models.DatabaseAccountUpdateParameters;
import com.azure.resourcemanager.cosmos.models.DefaultConsistencyLevel;
import com.azure.resourcemanager.cosmos.models.IpAddressOrRange;
import com.azure.resourcemanager.cosmos.models.ManagedServiceIdentity;
import com.azure.resourcemanager.cosmos.models.ManagedServiceIdentityUserAssignedIdentities;
import com.azure.resourcemanager.cosmos.models.MinimalTlsVersion;
import com.azure.resourcemanager.cosmos.models.NetworkAclBypass;
import com.azure.resourcemanager.cosmos.models.PeriodicModeBackupPolicy;
import com.azure.resourcemanager.cosmos.models.PeriodicModeProperties;
import com.azure.resourcemanager.cosmos.models.ResourceIdentityType;
import com.azure.resourcemanager.cosmos.models.VirtualNetworkRule;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for DatabaseAccounts Update.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2025-04-15/examples/
     * CosmosDBDatabaseAccountPatch.json
     */
    /**
     * Sample code: CosmosDBDatabaseAccountPatch.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBDatabaseAccountPatch(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getDatabaseAccounts().update("rg1", "ddb1",
            new DatabaseAccountUpdateParameters().withTags(mapOf("dept", "finance")).withLocation("westus")
                .withIdentity(new ManagedServiceIdentity().withType(ResourceIdentityType.SYSTEM_ASSIGNED_USER_ASSIGNED)
                    .withUserAssignedIdentities(mapOf(
                        "/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/eu2cgroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1",
                        new ManagedServiceIdentityUserAssignedIdentities())))
                .withConsistencyPolicy(
                    new ConsistencyPolicy().withDefaultConsistencyLevel(DefaultConsistencyLevel.BOUNDED_STALENESS)
                        .withMaxStalenessPrefix(200L).withMaxIntervalInSeconds(10))
                .withIpRules(Arrays.asList(new IpAddressOrRange().withIpAddressOrRange("23.43.230.120"),
                    new IpAddressOrRange().withIpAddressOrRange("110.12.240.0/12")))
                .withIsVirtualNetworkFilterEnabled(true)
                .withVirtualNetworkRules(Arrays.asList(new VirtualNetworkRule().withId(
                    "/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1")
                    .withIgnoreMissingVNetServiceEndpoint(false)))
                .withDefaultIdentity("FirstPartyIdentity").withEnableFreeTier(false).withEnableAnalyticalStorage(true)
                .withAnalyticalStorageConfiguration(
                    new AnalyticalStorageConfiguration().withSchemaType(AnalyticalStorageSchemaType.WELL_DEFINED))
                .withBackupPolicy(new PeriodicModeBackupPolicy().withPeriodicModeProperties(new PeriodicModeProperties()
                    .withBackupIntervalInMinutes(240).withBackupRetentionIntervalInHours(720)
                    .withBackupStorageRedundancy(BackupStorageRedundancy.LOCAL)))
                .withNetworkAclBypass(NetworkAclBypass.AZURE_SERVICES)
                .withNetworkAclBypassResourceIds(Arrays.asList(
                    "/subscriptions/subId/resourcegroups/rgName/providers/Microsoft.Synapse/workspaces/workspaceName"))
                .withCapacity(new Capacity().withTotalThroughputLimit(2000)).withEnablePartitionMerge(true)
                .withMinimalTlsVersion(MinimalTlsVersion.TLS).withEnableBurstCapacity(true)
                .withEnablePerRegionPerPartitionAutoscale(true),
            com.azure.core.util.Context.NONE);
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
