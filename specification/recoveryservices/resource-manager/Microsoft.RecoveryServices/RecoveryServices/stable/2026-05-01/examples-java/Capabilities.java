
import com.azure.resourcemanager.recoveryservices.models.CapabilitiesProperties;
import com.azure.resourcemanager.recoveryservices.models.DnsZone;
import com.azure.resourcemanager.recoveryservices.models.ResourceCapabilities;
import com.azure.resourcemanager.recoveryservices.models.VaultSubResourceType;
import java.util.Arrays;

/**
 * Samples for RecoveryServices Capabilities.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01/Capabilities.json
     */
    /**
     * Sample code: Capabilities for Microsoft.RecoveryServices/Vaults.
     * 
     * @param manager Entry point to RecoveryServicesManager.
     */
    public static void capabilitiesForMicrosoftRecoveryServicesVaults(
        com.azure.resourcemanager.recoveryservices.RecoveryServicesManager manager) {
        manager.recoveryServices().capabilitiesWithResponse("westus",
            new ResourceCapabilities().withType("Microsoft.RecoveryServices/Vaults")
                .withProperties(new CapabilitiesProperties()
                    .withDnsZones(Arrays.asList(new DnsZone().withSubResource(VaultSubResourceType.AZURE_BACKUP),
                        new DnsZone().withSubResource(VaultSubResourceType.AZURE_SITE_RECOVERY)))),
            com.azure.core.util.Context.NONE);
    }
}
