import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.DscpConfigurationInner;
import com.azure.resourcemanager.network.models.ProtocolType;
import com.azure.resourcemanager.network.models.QosDefinition;
import com.azure.resourcemanager.network.models.QosIpRange;
import com.azure.resourcemanager.network.models.QosPortRange;
import java.util.Arrays;

/** Samples for DscpConfiguration CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/DscpConfigurationCreate.json
     */
    /**
     * Sample code: Create DSCP Configuration.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createDSCPConfiguration(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getDscpConfigurations()
            .createOrUpdate(
                "rg1",
                "mydscpconfig",
                new DscpConfigurationInner()
                    .withLocation("eastus")
                    .withQosDefinitionCollection(
                        Arrays
                            .asList(
                                new QosDefinition()
                                    .withMarkings(Arrays.asList(1))
                                    .withSourceIpRanges(
                                        Arrays.asList(new QosIpRange().withStartIp("127.0.0.1").withEndIp("127.0.0.2")))
                                    .withDestinationIpRanges(
                                        Arrays
                                            .asList(new QosIpRange().withStartIp("127.0.10.1").withEndIp("127.0.10.2")))
                                    .withSourcePortRanges(
                                        Arrays
                                            .asList(
                                                new QosPortRange().withStart(10).withEnd(11),
                                                new QosPortRange().withStart(20).withEnd(21)))
                                    .withDestinationPortRanges(
                                        Arrays.asList(new QosPortRange().withStart(15).withEnd(15)))
                                    .withProtocol(ProtocolType.TCP),
                                new QosDefinition()
                                    .withMarkings(Arrays.asList(2))
                                    .withSourceIpRanges(
                                        Arrays.asList(new QosIpRange().withStartIp("12.0.0.1").withEndIp("12.0.0.2")))
                                    .withDestinationIpRanges(
                                        Arrays.asList(new QosIpRange().withStartIp("12.0.10.1").withEndIp("12.0.10.2")))
                                    .withSourcePortRanges(Arrays.asList(new QosPortRange().withStart(11).withEnd(12)))
                                    .withDestinationPortRanges(
                                        Arrays.asList(new QosPortRange().withStart(51).withEnd(52)))
                                    .withProtocol(ProtocolType.UDP))),
                Context.NONE);
    }
}
