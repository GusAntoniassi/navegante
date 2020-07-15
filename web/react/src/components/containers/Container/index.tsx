import * as React from 'react';
import {TiChevronLeft} from 'react-icons/ti';
import {AiOutlineSearch} from 'react-icons/ai';
import {Container as EntityContainer} from 'core';
import {ProgressBar} from 'components/ui/ProgressBar';

import 'stylesheets/06-components/container-item.scss';

type Props = {
    container: EntityContainer | null;
}

const collapseContainerAttr = (e: React.MouseEvent) => {
    const target: Element = e.target as Element;
    const containerAttr = target.closest('.container-attribute');

    if (containerAttr?.classList.contains('-collapsed')) {
        containerAttr.classList.remove('-collapsed');
    } else {
        containerAttr?.classList.add('-collapsed');
    }
};

const Chevron: React.FC = () => {
    return <span className="chevron" onClick={collapseContainerAttr}><TiChevronLeft className="icon"/></span>;
};

export const Container: React.FC<Props> = ({
                                               container
                                           }) => {
    if (container == null) return (null);
    let stats = container.statistics;

    return (
        <div className="container-item">
            <div className="heading">
                {container.name}
            </div>

            <ul className="content">
                <li className="container-attribute">ID: {container.id?.substr(0, 12)}</li>
                {container.ports?.length > 0 &&
                <li className="container-attribute -collapsed">
                    Ports <Chevron/>
                    <ul className="list">
                        {container.ports.map((port, index) => <li className="item" key={index}>- {port}</li>)}
                    </ul>
                </li>
                }

                {container.volumes?.length > 0 &&
                <li className="container-attribute -collapsed">
                    Volumes <Chevron/>
                    <ul className="list">
                        {container.volumes.map((volume, index) => <li className="item" key={index}>- {volume}</li>)}
                    </ul>
                </li>
                }

                {container.networks?.length > 0 &&
                <li className="container-attribute -collapsed">
                    Networks <Chevron/>
                    <ul className="list">
                        {container.networks.map((network, index) => <li className="item" key={index}>- {network}</li>)}
                    </ul>
                </li>
                }

                <li className="container-attribute -collapsed">
                    Stats <Chevron/>
                    <ul className="list">
                        <li className="item display-flex">
                            <span style={{marginRight: "10px"}}>CPU:</span>
                            <ProgressBar inline={true} progress={stats.cpuPercent}></ProgressBar>
                        </li>
                        <li className="item display-flex">
                            <span style={{marginRight: "10px"}}>Mem:</span>
                            <ProgressBar inline={true} progress={stats.memoryPercent}></ProgressBar>
                        </li>
                        <li className="item">Mem usg: {stats.memoryUsage.toHuman()} / {stats.memoryTotal.toHuman()}</li>
                        <li className="item">Net
                            I/O: {stats.networkInput.toHuman()} / {stats.networkOutput.toHuman()}</li>
                        <li className="item">Block I/O: {stats.blockRead.toHuman()} / {stats.blockWrite.toHuman()}</li>
                    </ul>
                </li>
            </ul>

            <div className="footer text-right">
                <AiOutlineSearch className="cursor-pointer"/>
            </div>
        </div>
    );
};
