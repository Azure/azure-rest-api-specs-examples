
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
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
     * x-ms-original-file: 2025-10-01-preview/NodeTypePutOperationStateless_example.json
     */
    /**
     * Sample code: Put an stateless node type with temporary disk for service fabric.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void putAnStatelessNodeTypeWithTemporaryDiskForServiceFabric(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager)
        throws IOException {
        manager.nodeTypes().define("BE").withExistingManagedCluster("resRg", "myCluster").withIsPrimary(false)
            .withVmInstanceCount(10).withVmSize("Standard_DS3").withVmImagePublisher("MicrosoftWindowsServer")
            .withVmImageOffer("WindowsServer").withVmImageSku("2016-Datacenter-Server-Core")
            .withVmImageVersion("latest")
            .withVmExtensions(Arrays.asList(new VmssExtension().withName("Microsoft.Azure.Geneva.GenevaMonitoring")
                .withPublisher("Microsoft.Azure.Geneva").withType("GenevaMonitoring").withTypeHandlerVersion("2.0")
                .withAutoUpgradeMinorVersion(true)
                .withSettings(SerializerFactory.createDefaultManagementSerializerAdapter().deserialize("{}",
                    Object.class, SerializerEncoding.JSON))))
            .withIsStateless(true).withMultiplePlacementGroups(true).withEnableEncryptionAtHost(true)
            .withUseTempDataDisk(true).create();
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
