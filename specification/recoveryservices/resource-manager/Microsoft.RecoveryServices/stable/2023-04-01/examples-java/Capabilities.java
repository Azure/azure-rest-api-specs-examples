import com.azure.resourcemanager.recoveryservices.models.CapabilitiesProperties;
import com.azure.resourcemanager.recoveryservices.models.DnsZone;
import com.azure.resourcemanager.recoveryservices.models.ResourceCapabilities;
import com.azure.resourcemanager.recoveryservices.models.VaultSubResourceType;
import java.util.Arrays;

/** Samples for RecoveryServices Capabilities. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2023-04-01/examples/Capabilities.json
     */
    /**
     * Sample code: Capabilities for Microsoft.RecoveryServices/Vaults.
     *
     * @param manager Entry point to RecoveryServicesManager.
     */
    public static void capabilitiesForMicrosoftRecoveryServicesVaults(
        com.azure.resourcemanager.recoveryservices.RecoveryServicesManager manager) {
        manager
            .recoveryServices()
            .capabilitiesWithResponse(
                "westus",
                new ResourceCapabilities()
                    .withType("Microsoft.RecoveryServices/Vaults")
                    .withProperties(
                        new CapabilitiesProperties()
                            .withDnsZones(
                                Arrays
                                    .asList(
                                        new DnsZone().withSubResource(VaultSubResourceType.AZURE_BACKUP),
                                        new DnsZone().withSubResource(VaultSubResourceType.AZURE_SITE_RECOVERY)))),
                com.azure.core.util.Context.NONE);
    }
}
