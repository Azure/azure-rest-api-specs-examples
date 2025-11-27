
import com.azure.core.management.SubResource;
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.servicefabricmanagedclusters.models.DiskType;
import com.azure.resourcemanager.servicefabricmanagedclusters.models.VaultCertificate;
import com.azure.resourcemanager.servicefabricmanagedclusters.models.VaultSecretGroup;
import com.azure.resourcemanager.servicefabricmanagedclusters.models.VmManagedIdentity;
import com.azure.resourcemanager.servicefabricmanagedclusters.models.VmssExtension;
import java.io.IOException;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for NodeTypes CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/NodeTypePutOperationAutoScale_example.json
     */
    /**
     * Sample code: Put a node type with auto-scale parameters.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void putANodeTypeWithAutoScaleParameters(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager)
        throws IOException {
        manager.nodeTypes().define("BE").withExistingManagedCluster("resRg", "myCluster").withIsPrimary(false)
            .withVmInstanceCount(-1).withDataDiskSizeGB(200).withDataDiskType(DiskType.PREMIUM_LRS)
            .withPlacementProperties(mapOf("HasSSD", "true", "NodeColor", "green", "SomeProperty", "5"))
            .withCapacities(mapOf("ClientConnections", "65536")).withVmSize("Standard_DS3")
            .withVmImagePublisher("MicrosoftWindowsServer").withVmImageOffer("WindowsServer")
            .withVmImageSku("2016-Datacenter-Server-Core").withVmImageVersion("latest")
            .withVmSecrets(Arrays.asList(new VaultSecretGroup().withSourceVault(new SubResource().withId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.KeyVault/vaults/myVault"))
                .withVaultCertificates(Arrays.asList(new VaultCertificate()
                    .withCertificateUrl(
                        "https://myVault.vault.azure.net:443/secrets/myCert/ef1a31d39e1f46bca33def54b6cda54c")
                    .withCertificateStore("My")))))
            .withVmExtensions(Arrays.asList(new VmssExtension().withName("Microsoft.Azure.Geneva.GenevaMonitoring")
                .withPublisher("Microsoft.Azure.Geneva").withType("GenevaMonitoring").withTypeHandlerVersion("2.0")
                .withAutoUpgradeMinorVersion(true)
                .withSettings(SerializerFactory.createDefaultManagementSerializerAdapter().deserialize("{}",
                    Object.class, SerializerEncoding.JSON))))
            .withVmManagedIdentity(new VmManagedIdentity().withUserAssignedIdentities(Arrays.asList(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myIdentity",
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myIdentity2")))
            .withIsStateless(true).withMultiplePlacementGroups(true).create();
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
