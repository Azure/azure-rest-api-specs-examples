```java
import com.azure.resourcemanager.mobilenetwork.models.AttachedDataNetworkResourceId;
import com.azure.resourcemanager.mobilenetwork.models.MobileNetworkResourceId;
import com.azure.resourcemanager.mobilenetwork.models.SimPolicyResourceId;
import com.azure.resourcemanager.mobilenetwork.models.SimStaticIpProperties;
import com.azure.resourcemanager.mobilenetwork.models.SimStaticIpPropertiesStaticIp;
import com.azure.resourcemanager.mobilenetwork.models.SliceResourceId;
import java.util.Arrays;

/** Samples for Sims CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-01-01-preview/examples/SimCreate.json
     */
    /**
     * Sample code: Create sim.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void createSim(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager
            .sims()
            .define("testSim")
            .withRegion("testLocation")
            .withExistingResourceGroup("rg1")
            .withInternationalMobileSubscriberIdentity("00000")
            .withIntegratedCircuitCardIdentifier("8900000000000000000")
            .withAuthenticationKey("00000000000000000000000000000000")
            .withOperatorKeyCode("00000000000000000000000000000000")
            .withMobileNetwork(
                new MobileNetworkResourceId()
                    .withId(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork"))
            .withDeviceType("Video camera")
            .withSimPolicy(
                new SimPolicyResourceId()
                    .withId(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/simPolicies/MySimPolicy"))
            .withStaticIpConfiguration(
                Arrays
                    .asList(
                        new SimStaticIpProperties()
                            .withAttachedDataNetwork(
                                new AttachedDataNetworkResourceId()
                                    .withId(
                                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/TestPacketCoreCP/packetCoreDataPlanes/TestPacketCoreDP/attachedDataNetworks/TestAttachedDataNetwork"))
                            .withSlice(
                                new SliceResourceId()
                                    .withId(
                                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/slices/testSlice"))
                            .withStaticIp(new SimStaticIpPropertiesStaticIp().withIpv4Address("2.4.0.1"))))
            .create();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mobilenetwork_1.0.0-beta.1/sdk/mobilenetwork/azure-resourcemanager-mobilenetwork/README.md) on how to add the SDK to your project and authenticate.
