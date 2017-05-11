interface Pipeline {
    Name: string;
    Stages: Stage[];
}

interface Stage {
    Name: string;
    Jobs: Job[];
}

interface Job {
    Name: string;
    Steps: Step[];
}

interface Step {
    OridinalPosition: number;
}