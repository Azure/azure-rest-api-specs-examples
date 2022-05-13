Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-scvmm_1.0.0-beta.1/sdk/scvmm/azure-resourcemanager-scvmm/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.scvmm.models.AllocationMethod;
import com.azure.resourcemanager.scvmm.models.HardwareProfileUpdate;
import com.azure.resourcemanager.scvmm.models.NetworkInterfacesUpdate;
import com.azure.resourcemanager.scvmm.models.NetworkProfileUpdate;
import com.azure.resourcemanager.scvmm.models.StorageProfileUpdate;
import com.azure.resourcemanager.scvmm.models.VirtualDiskUpdate;
import com.azure.resourcemanager.scvmm.models.VirtualMachine;
import com.azure.resourcemanager.scvmm.models.VirtualMachineUpdateProperties;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for VirtualMachines Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/UpdateVirtualMachine.json
     */
    /**
     * Sample code: UpdateVirtualMachine.
     *
     * @param manager Entry point to ScvmmManager.
     */
    public static void updateVirtualMachine(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        VirtualMachine resource =
            manager.virtualMachines().getByResourceGroupWithResponse("testrg", "DemoVM", Context.NONE).getValue();
        resource
            .update()
            .withTags(mapOf("tag1", "value1", "tag2", "value2"))
            .withProperties(
                new VirtualMachineUpdateProperties()
                    .withHardwareProfile(new HardwareProfileUpdate().withMemoryMB(4096).withCpuCount(4))
                    .withStorageProfile(
                        new StorageProfileUpdate()
                            .withDisks(Arrays.asList(new VirtualDiskUpdate().withName("test").withDiskSizeGB(10))))
                    .withNetworkProfile(
                        new NetworkProfileUpdate()
                            .withNetworkInterfaces(
                                Arrays
                                    .asList(
                                        new NetworkInterfacesUpdate()
                                            .withName("test")
                                            .withIpv4AddressType(AllocationMethod.DYNAMIC)
                                            .withIpv6AddressType(AllocationMethod.DYNAMIC)
                                            .withMacAddressType(AllocationMethod.STATIC)))))
            .apply();
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
```
