
import com.azure.core.management.SubResource;
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.servicefabricmanagedclusters.models.AdditionalNetworkInterfaceConfiguration;
import com.azure.resourcemanager.servicefabricmanagedclusters.models.DiskType;
import com.azure.resourcemanager.servicefabricmanagedclusters.models.EvictionPolicyType;
import com.azure.resourcemanager.servicefabricmanagedclusters.models.FrontendConfiguration;
import com.azure.resourcemanager.servicefabricmanagedclusters.models.IpConfiguration;
import com.azure.resourcemanager.servicefabricmanagedclusters.models.IpConfigurationPublicIpAddressConfiguration;
import com.azure.resourcemanager.servicefabricmanagedclusters.models.IpTag;
import com.azure.resourcemanager.servicefabricmanagedclusters.models.PrivateIpAddressVersion;
import com.azure.resourcemanager.servicefabricmanagedclusters.models.PublicIpAddressVersion;
import com.azure.resourcemanager.servicefabricmanagedclusters.models.SecurityEncryptionType;
import com.azure.resourcemanager.servicefabricmanagedclusters.models.SecurityType;
import com.azure.resourcemanager.servicefabricmanagedclusters.models.VaultCertificate;
import com.azure.resourcemanager.servicefabricmanagedclusters.models.VaultSecretGroup;
import com.azure.resourcemanager.servicefabricmanagedclusters.models.VmApplication;
import com.azure.resourcemanager.servicefabricmanagedclusters.models.VmManagedIdentity;
import com.azure.resourcemanager.servicefabricmanagedclusters.models.VmSetupAction;
import com.azure.resourcemanager.servicefabricmanagedclusters.models.VmssDataDisk;
import com.azure.resourcemanager.servicefabricmanagedclusters.models.VmssExtension;
import com.azure.resourcemanager.servicefabricmanagedclusters.models.VmssExtensionSetupOrder;
import java.io.IOException;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for NodeTypes CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/NodeTypePutOperation_example_max.json
     */
    /**
     * Sample code: Put a node type with maximum parameters.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void putANodeTypeWithMaximumParameters(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager)
        throws IOException {
        manager.nodeTypes().define("BE-testResourceGroup-testRegion-test")
            .withExistingManagedCluster("resRg", "myCluster").withIsPrimary(false).withVmInstanceCount(10)
            .withDataDiskSizeGB(200).withDataDiskType(DiskType.PREMIUM_LRS).withDataDiskLetter("S")
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
                    Object.class, SerializerEncoding.JSON))
                .withForceUpdateTag("v.1.0").withEnableAutomaticUpgrade(true)
                .withSetupOrder(Arrays.asList(VmssExtensionSetupOrder.BEFORE_SFRUNTIME))))
            .withVmManagedIdentity(new VmManagedIdentity().withUserAssignedIdentities(Arrays.asList(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myIdentity",
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myIdentity2")))
            .withIsStateless(true).withMultiplePlacementGroups(true)
            .withFrontendConfigurations(Arrays.asList(new FrontendConfiguration().withLoadBalancerBackendAddressPoolId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/loadBalancers/test-LB/backendAddressPools/LoadBalancerBEAddressPool")
                .withLoadBalancerInboundNatPoolId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/loadBalancers/test-LB/inboundNatPools/LoadBalancerNATPool")
                .withApplicationGatewayBackendAddressPoolId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/applicationGateways/appgw-test/backendAddressPools/appgwBepoolTest")))
            .withAdditionalDataDisks(Arrays.asList(
                new VmssDataDisk().withLun(1).withDiskSizeGB(256).withDiskType(DiskType.STANDARD_SSD_LRS)
                    .withDiskLetter("F"),
                new VmssDataDisk().withLun(2).withDiskSizeGB(150).withDiskType(DiskType.PREMIUM_LRS)
                    .withDiskLetter("G")))
            .withEnableEncryptionAtHost(true).withEnableAcceleratedNetworking(true)
            .withUseDefaultPublicLoadBalancer(true).withEnableOverProvisioning(false).withIsSpotVM(true)
            .withUseEphemeralOSDisk(true).withSpotRestoreTimeout("PT30M")
            .withEvictionPolicy(EvictionPolicyType.DEALLOCATE)
            .withSubnetId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1")
            .withVmSetupActions(Arrays.asList(VmSetupAction.ENABLE_CONTAINERS, VmSetupAction.ENABLE_HYPERV))
            .withSecurityType(SecurityType.CONFIDENTIAL_VM)
            .withSecurityEncryptionType(SecurityEncryptionType.DISK_WITH_VMGUEST_STATE).withSecureBootEnabled(true)
            .withEnableNodePublicIp(true).withEnableNodePublicIPv6(true)
            .withNatGatewayId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/natGateways/myNatGateway")
            .withServiceArtifactReferenceId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Compute/galleries/myGallery/serviceArtifacts/myServiceArtifact/vmArtifactsProfiles/myVmArtifactProfile")
            .withDscpConfigurationId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/dscpConfigurations/myDscpConfig")
            .withAdditionalNetworkInterfaceConfigurations(Arrays.asList(new AdditionalNetworkInterfaceConfiguration()
                .withName("nic-1").withEnableAcceleratedNetworking(true)
                .withDscpConfiguration(new SubResource().withId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/dscpConfigurations/myDscpConfig"))
                .withIpConfigurations(Arrays.asList(new IpConfiguration().withName("ipconfig-1")
                    .withApplicationGatewayBackendAddressPools(Arrays.asList(new SubResource().withId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/applicationGateways/appgw-test/backendAddressPools/appgwBepoolTest")))
                    .withLoadBalancerBackendAddressPools(Arrays.asList(new SubResource().withId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/loadBalancers/test-LB/backendAddressPools/LoadBalancerBEAddressPool")))
                    .withLoadBalancerInboundNatPools(Arrays.asList(new SubResource().withId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/loadBalancers/test-LB/inboundNatPools/LoadBalancerNATPool")))
                    .withSubnet(new SubResource().withId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"))
                    .withPrivateIpAddressVersion(PrivateIpAddressVersion.IPV4)
                    .withPublicIpAddressConfiguration(new IpConfigurationPublicIpAddressConfiguration()
                        .withName("publicip-1")
                        .withIpTags(Arrays.asList(new IpTag().withIpTagType("RoutingPreference").withTag("Internet")))
                        .withPublicIpAddressVersion(PublicIpAddressVersion.IPV4))))))
            .withComputerNamePrefix("BE")
            .withVmApplications(Arrays.asList(new VmApplication()
                .withConfigurationReference("https://mystorageaccount.blob.core.windows.net/containername/blobname")
                .withEnableAutomaticUpgrade(true).withOrder(1)
                .withPackageReferenceId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Compute/galleries/myGallery/applications/myApplication/versions/1.0.0")
                .withVmGalleryTags("{\"Tag1\":\"Value1\",\"Tag2\":\"Value2\"}")
                .withTreatFailureAsDeploymentFailure(false)))
            .withIsOutboundOnly(true).create();
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
