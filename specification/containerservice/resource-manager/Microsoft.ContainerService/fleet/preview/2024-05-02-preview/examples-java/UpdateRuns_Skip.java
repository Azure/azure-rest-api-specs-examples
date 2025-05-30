
import com.azure.resourcemanager.containerservicefleet.models.SkipProperties;
import com.azure.resourcemanager.containerservicefleet.models.SkipTarget;
import com.azure.resourcemanager.containerservicefleet.models.TargetType;
import java.util.Arrays;

/**
 * Samples for UpdateRuns Skip.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/preview/2024-05-02-preview/
     * examples/UpdateRuns_Skip.json
     */
    /**
     * Sample code: Skips one or more member/group/stage/afterStageWait(s) of an UpdateRun.
     * 
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void skipsOneOrMoreMemberGroupStageAfterStageWaitSOfAnUpdateRun(
        com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.updateRuns().skip("rg1", "fleet1", "run1",
            new SkipProperties()
                .withTargets(Arrays.asList(new SkipTarget().withType(TargetType.MEMBER).withName("member-one"),
                    new SkipTarget().withType(TargetType.AFTER_STAGE_WAIT).withName("stage1"))),
            null, com.azure.core.util.Context.NONE);
    }
}
