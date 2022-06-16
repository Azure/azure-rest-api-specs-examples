import com.azure.resourcemanager.mobilenetwork.models.Ambr;
import com.azure.resourcemanager.mobilenetwork.models.PccRuleConfiguration;
import com.azure.resourcemanager.mobilenetwork.models.PccRuleQosPolicy;
import com.azure.resourcemanager.mobilenetwork.models.PreemptionCapability;
import com.azure.resourcemanager.mobilenetwork.models.PreemptionVulnerability;
import com.azure.resourcemanager.mobilenetwork.models.QosPolicy;
import com.azure.resourcemanager.mobilenetwork.models.SdfDirection;
import com.azure.resourcemanager.mobilenetwork.models.ServiceDataFlowTemplate;
import com.azure.resourcemanager.mobilenetwork.models.TrafficControlPermission;
import java.util.Arrays;

/** Samples for Services CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-01-01-preview/examples/ServiceCreate.json
     */
    /**
     * Sample code: Create service.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void createService(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager
            .services()
            .define("TestService")
            .withRegion("eastus")
            .withExistingMobileNetwork("rg1", "testMobileNetwork")
            .withServicePrecedence(255)
            .withPccRules(
                Arrays
                    .asList(
                        new PccRuleConfiguration()
                            .withRuleName("default-rule")
                            .withRulePrecedence(255)
                            .withRuleQosPolicy(
                                new PccRuleQosPolicy()
                                    .withFiveQi(9)
                                    .withAllocationAndRetentionPriorityLevel(9)
                                    .withPreemptionCapability(PreemptionCapability.NOT_PREEMPT)
                                    .withPreemptionVulnerability(PreemptionVulnerability.PREEMPTABLE)
                                    .withMaximumBitRate(new Ambr().withUplink("500 Mbps").withDownlink("1 Gbps")))
                            .withTrafficControl(TrafficControlPermission.ENABLED)
                            .withServiceDataFlowTemplates(
                                Arrays
                                    .asList(
                                        new ServiceDataFlowTemplate()
                                            .withTemplateName("IP-to-server")
                                            .withDirection(SdfDirection.UPLINK)
                                            .withProtocol(Arrays.asList("ip"))
                                            .withRemoteIpList(Arrays.asList("10.3.4.0/24"))
                                            .withPorts(Arrays.asList())))))
            .withServiceQosPolicy(
                new QosPolicy()
                    .withFiveQi(9)
                    .withAllocationAndRetentionPriorityLevel(9)
                    .withPreemptionCapability(PreemptionCapability.NOT_PREEMPT)
                    .withPreemptionVulnerability(PreemptionVulnerability.PREEMPTABLE)
                    .withMaximumBitRate(new Ambr().withUplink("500 Mbps").withDownlink("1 Gbps")))
            .create();
    }
}
